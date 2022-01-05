package command
/*
    Overview:
        Ordinary statements are perishable
            - cannot undo field assignment
            - cannot directly serialize a sequence of actions (calls)
        Want an object that represents an operation
    DoCommand: An object which represents an instruction to perform a particular action. Contains all information necessary for the action to be taken.

    Summary:
        Encapsulate all details of an operation in a separate object
        Define functions for applying the command (either in the command itself, or elsewhere)
        Optionally define instructions for undoing the command
        Can create composite commands (a.k.a. macros)
 */
