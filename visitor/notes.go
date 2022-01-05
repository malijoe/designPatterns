package visitor

/*
    Overview:
        Need to define a new operation on an entire type hierarchy
        Do not want to keep modifying every type in the hierarchy
        Want to have the new functionality separate (SRP)
        This approach is often used for traversal
            Alternative to Iterator
            Hierarchy members help you traverse themselves

    Visitor: A pattern where a component (visitor) is allowed to traverse the entire hierarchy of types. Implemented by propagating a single Accept() method throughout the entire hierarchy.

    Dispatch:
        Which function to call?
        Single dispatch: depends on name of request and type of receiver
        Double dispatch: depends on name of request and type of two receivers (type of visitor, type of element being visited)

    Summary:
        Propagate an Accept(v *Visitor) method throughout the entire hierarchy
        Create a visitor with VisitFoo(f Foo), VisitBar(b Bar), ... for each element in the hierarchy
        Each Accept() simply calls Visitor.VisitXxx(self)
*/
