package state
/*
    Overview:
        Changes in state can be explicit or in response to an event (Observer pattern)

    State: A pattern in which the object's behavior is determined by its state. An object transitions from one state to another (something needs to trigger a transition).
            A formalized construct which manages state and transitions is called a state machine.

    Summary:
        Given sufficient complexity, it pays to formally define possible states and events/triggers
        Can define:
            State entry/exit behaviors
            Action when a particular event causes a transition
            Guard conditions enabling/disabling a transition
            Default action when no transitions are found for an event
 */