package flyweight

/*
    Avoid redundancy when storing data
    e.g. mmorpg
        plenty of users with identical first/last names
        no sens in storing same first/last name over and over again
        store a list of names and references to them (indices, pointers, etc.)
    e.g. bold or italic text formatting
        Don't want each character to have a formatting character
        Operate on ranges (e.g. line number, start/end positions)

    Flyweight: A space optimization technique that lets us use less memory by storing externally the data associated with similar objects.

    Summary:
        Store common data externally
        Specify an index or a pointer into the external data store
        Define the idea of 'ranges' on homogeneous collections and store data related to those ranges
 */