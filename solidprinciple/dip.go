package solidprinciple

/*
	Dependency Inversion Principle
	HLM should not depend on LLM
	Both should depend on abstraction
*/

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

type Relationships struct {
	relations []Info
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
		}
		result = append(result, r.relations[i].to)
	}
	return result
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

type RelationShipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

type Research struct {
	//relationships Relationships
	browser RelationShipBrowser
}

//func (r *Research) Investigate() {
//	relations := r.relationships.relations
//	for _, rel := range relations {
//		if rel.from.name == "John" && rel.relationship == Parent {
//			fmt.Println("John has a child called", rel.to.name)
//		}
//	}
//}
