package memento
/*
    Overview:
        An object or system goes through changes
        There are different ways of navigating those changes
        One way os to record every change (Command) and teach a command to 'undo' itself
            Also apart of CQRS = Command Query Responsibility Segregation
        Another it to simply save snapshots of the system
    Memento: A token representing the system state. Lets us roll back to the state when the token was generated. May or may not directly expose state information.

    Memento vs Flyweight
        Both patterns provide a 'token' clients can hold on to
        Memento is used only to be fed back into the system
            No public/mutable state
            No methods
        A flyweight is similar to an ordinary reference to an object
            Can mutate state
            Can provide additional functionality (fields/methods)

    Summary:
        Mementos are used to roll back states arbitrarily
        A memento is simply a token/handler with (typically) no methods of its own
        A memento is not required to expose directly the state(s) to which it reverts to the system
        Can be used to implement undo/redo
 */
