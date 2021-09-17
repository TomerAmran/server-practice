package cache

type contextKeyType int

const (
	ExecutorContetKey contextKeyType = iota
)

//it allows us to write some other implementatin of the cache so we can test it?
type Executer interface {
	Get(key string) ([]byte, error)
	Set(key string, data []byte) error
	DeleteAll() error
}

type InMemory struct {
	data map[string][]byte
}

func (i *InMemory) Get(key string) ([]byte, error){
	return i.data[key],nil
}

func (i *InMemory) Set(key string, data []byte) error{
	i.data[key] = data
	return nil
}

func (i *InMemory) DeleteAll() error{
	for k := range i.data {
		delete(i.data, k)
	}
	return nil
}

func(i *InMemory) Delete(key string) error{
	delete(i.data,key)
	return nil
}

func NewInMemoryCache(/*config*/)*InMemory {
	return &InMemory{
		make(map[string][]byte),
	}
}