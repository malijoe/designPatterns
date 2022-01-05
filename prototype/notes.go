package prototype

/*
    Overview:
        Complicated objects (e.g. cars) aren't designed from scratch
            they reiterate existing designs
        An existing (partially or fully constructed) design is a Prototype
        We make a copy of the prototype and customize it
            requires 'deep copy' support (make sure we are copying pointer values not addresses)
        We make the cloning convenient (Factory)

    Prototype - A partially or fully initialized object that you copy (clone) and make use of
 */
/*
    Summary:
        To implement a prototype, partially construct an object and store it somewhere
        Deep copy the prototype
        Customize the resulting instance
        A prototype factory provides a convenient API for using prototypes
 */