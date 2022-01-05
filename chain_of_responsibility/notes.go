package chain_of_responsibility

/*
    Chain or Responsibility: A chain of components who all get a chance to process a command or a query, optionally having a default processing implementation and an ability to terminate the processing chain.

    Command Query Separation:
        Command: asking for an action or change (mutable operation: changes state)
        Query: asking for information (immutable operation: doesn't change state)
    CQS: having separate means of sending commands and queries to direct field access.

    Summary:
        Chain of Responsibility can be implemented as a linked list of pointers or a centralized construct
        Enlist objects in the chain, possibly controlling their order
        Control object removal from chain
 */
