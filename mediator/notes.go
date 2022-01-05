package mediator
 /*
    Overview:
        Components may go in and out of a system at any time
            - Chat room participants
            - players in an mmorpg
        It makes no sense for them to have direct references (pointers to one another
            Those references may go dead
        Solution: have them all refer to some central component that facilitates communication

    Mediator: A component that facilitates communication between other components without them necessarily being aware of each other or having direct (references) access to each other.

    Summary:
        Create the mediator and have each object in the system points to it
            e.g., assign a field in factory function
        Mediator engages in bidirectional communication with its connected components
        Mediate has methods that the components can call
        Components have methods the mediator can call
        Event processing (e.g., Rx) libraries make communication easier to implement
  */