package handlers

import "sync"

type inMemoryMap struct {
	Mu    sync.Mutex
	Store map[string]string
}

func (m inMemoryMap) Write(data secretData) {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	m.Store[data.Id] = data.PlainText

}

func (m inMemoryMap) Read(id string) string {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	data := m.Store[id]
	delete(m.Store, id)

	return data
}
