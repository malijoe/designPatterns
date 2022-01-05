package adapter
/*
   Overview: adapters give us the interface we require from the interface we have

   Adapter: A construct which adapts an existing Interface X to conform to the required interface Y.
*/

/*
    Summary:
        Implementing an Adapter is easy
        Determine the API you have and the API you need
        Create a component which aggregates (has a pointer to, ...) the adaptee
        Intermediate representations can pile up: use caching and other optimizations
 */