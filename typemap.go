package main

import (
	"fmt"
	"reflect"
)

type StructField4JSON struct {
	Original reflect.StructField `json:"-"`
	Name     string
	PkgPath  string
	TypeFqn  string
	Tag      reflect.StructTag
}

type Type4JSON struct {
	Original reflect.Type `json:"-"`
	Name     string
	PkgPath  string
	Size     uintptr
	String   string
	Kind     reflect.Kind
	KindName string
	ElemFqn  string
	KeyFqn   string
	Fields   []*StructField4JSON
}

type TypeMap map[string]*Type4JSON

func (m TypeMap) Start(targets ...interface{}) {
	for _, target := range targets {
		t := reflect.TypeOf(target)
		m.Walk(t)
	}
}

func (m TypeMap) Walk(t reflect.Type) string {
	fqn := m.genFqnFor(t)
	if _, ok := m[fqn]; ok {
		return fqn
	}

	r := &Type4JSON{
		Original: t,
		Name:     t.Name(),
		PkgPath:  t.PkgPath(),
		Size:     t.Size(),
		String:   t.String(),
		Kind:     t.Kind(),
		KindName: t.Kind().String(),
	}
	m[fqn] = r

	switch t.Kind() {
	case
		reflect.Array,
		reflect.Slice,
		reflect.Chan,
		reflect.Ptr:
		r.ElemFqn = m.Walk(t.Elem())

	case reflect.Map:
		r.KeyFqn = m.Walk(t.Key())
		r.ElemFqn = m.Walk(t.Elem())

	case reflect.Struct:
		num := t.NumField()
		fields := make([]*StructField4JSON, num)
		for i := 0; i < num; i++ {
			f := t.Field(i)
			fields[i] = &StructField4JSON{
				Original: f,
				Name:     f.Name,
				PkgPath:  f.PkgPath,
				TypeFqn:  m.Walk(t.Field(i).Type),
				Tag:      f.Tag,
			}
		}
		r.Fields = fields
	}

	return fqn
}

func (m TypeMap) genFqnFor(t reflect.Type) string {
	pkg := t.PkgPath()
	name := t.Name()
	if name == "" {
		switch t.Kind() {
		case reflect.Array:
			name = fmt.Sprintf("[%d]%s", t.Len(), m.genFqnFor(t.Elem()))

		case reflect.Slice:
			name = fmt.Sprintf("[]%s", m.genFqnFor(t.Elem()))

		case reflect.Chan:
			name = fmt.Sprintf("ch:%s", m.genFqnFor(t.Elem()))

		case reflect.Ptr:
			name = fmt.Sprintf("*%s", m.genFqnFor(t.Elem()))

		case reflect.Map:
			name = fmt.Sprintf("map[%s]%s", m.genFqnFor(t.Key()), m.genFqnFor(t.Elem()))

		default:
			name = t.String()
			if name == "" {
				name = fmt.Sprintf("(anonymous %s)", t.Kind().String())
			}
		}
	}

	if pkg == "" {
		return name
	} else {
		return fmt.Sprintf("%s.%s", pkg, name)
	}
}
