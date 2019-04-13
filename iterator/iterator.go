//  Iterator 迭代器模式：
//        提供一种方法顺序访问一个聚合对象中的各个元素，而又不暴露该对象的内部表示
package iterator

type Book struct {
	name string
}

type Iterator interface {
	first() interface{}
	next() interface{}
}

type BookGroup struct {
	books []Book
}

func (b *BookGroup) add(newbook Book) {
	if b == nil {
		return
	}
	b.books = append(b.books, newbook)
}

type BookIterator struct {
	g *BookGroup
	index int
}
func (b *BookGroup) createIterator() *BookIterator {
	if b == nil {
		return nil
	}
	return &BookIterator{b, 0}

}

func (b *BookIterator) first() interface{} {
	if b == nil {
		return nil
	}
	if len(b.g.books) > 0 {
		b.index = 0
		return b.g.books[b.index]
	}
	return nil
}

func (b *BookIterator) next() interface{} {
	if b == nil {
		return nil
	}
	if len(b.g.books) > b.index + 1 {
		b.index++
		return b.g.books[b.index]
	}
	return nil
}
