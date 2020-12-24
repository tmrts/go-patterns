The [Memento pattern](https://en.wikipedia.org/wiki/Memento_pattern) allows the state of an object to be saved and restored. It helps us write undo-redo operations through the use of the three componants: 

* Originator: It is the actual object whose state is saved as a memento. 
* Memento: This is the object which saves the state of the originator
* Caretaker: This is the object that saves multiple mementos. Given an index, it returns the corresponding memento. 

The originator defines two methods. savememento() and restorememento()

* savememento()- in this method the originator saves its internal state into a memento object.
* restorememento()- this method takes input as a memento object. The originator restores itself to the pass memento.  Hence a previous state is restored.