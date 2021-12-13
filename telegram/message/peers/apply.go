package peers

import (
	"context"

	"github.com/go-faster/errors"
	"go.uber.org/multierr"

	"github.com/gotd/td/tg"
)

func (m *Manager) applyUsers(ctx context.Context, users ...tg.UserClass) error {
	return saveUsers(ctx, m.storage, users...)
}

func (m *Manager) applyChats(ctx context.Context, chats ...tg.ChatClass) error {
	return saveChats(ctx, m.storage, chats...)
}

func (m *Manager) applyEntities(ctx context.Context, users []tg.UserClass, chats []tg.ChatClass) error {
	return multierr.Append(m.applyUsers(ctx, users...), m.applyChats(ctx, chats...))
}

func (m *Manager) updateContacts(ctx context.Context) ([]tg.UserClass, error) {
	if err := m.phone.Acquire(ctx, 1); err != nil {
		return nil, errors.Wrap(err, "acquire phone")
	}
	defer m.phone.Release(1)

	hash, err := m.storage.GetContactsHash(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get contacts hash")
	}

	r, err := m.api.ContactsGetContacts(ctx, hash)
	if err != nil {
		return nil, errors.Wrap(err, "get contacts")
	}

	switch c := r.(type) {
	case *tg.ContactsContacts:
		if err := m.applyUsers(ctx, c.Users...); err != nil {
			return nil, errors.Wrap(err, "update users")
		}

		me, ok := m.me.Load()
		if !ok {
			return nil, nil
		}

		if err := m.storage.SaveContactsHash(ctx, contactsHash(me.ID, c)); err != nil {
			return nil, errors.Wrap(err, "update contacts hash")
		}
		return c.Users, nil
	case *tg.ContactsContactsNotModified:
		return nil, nil
	default:
		return nil, errors.Errorf("unexpected type %T", r)
	}
}

type vectorHash struct {
	state uint64
}

// See https://github.com/tdlib/td/blob/aa8a4979df8fc56032f134471a2cb939a7b0839f/td/telegram/misc.cpp#L242.
func (h *vectorHash) apply(n uint64) {
	h.state ^= h.state >> 21
	h.state ^= h.state << 35
	h.state ^= h.state >> 4
	h.state += n
}

// See https://github.com/tdlib/td/blob/aa8a4979df8fc56032f134471a2cb939a7b0839f/td/telegram/ContactsManager.cpp#L5125.
func contactsHash(myID int64, contacts *tg.ContactsContacts) int64 {
	contacts.MapUsers().SortStableByID()

	var lesserIDx = len(contacts.Users) - 1
	for i, user := range contacts.Users {
		if user.GetID() < myID {
			lesserIDx = i
			break
		}
	}

	var h vectorHash
	h.apply(uint64(len(contacts.Users)))
	for i, contact := range contacts.Users {
		h.apply(uint64(contact.GetID()))
		if i == lesserIDx {
			h.apply(uint64(myID))
		}
	}

	return int64(h.state)
}
