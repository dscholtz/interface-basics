# interface-basics
basics of GO interfaces

The area project teaches the following:
An interface is a contract.
If an interface called Shape has one function: Area() float64,
then any type (like Circle, Square, etc.) that has a method Area() float64 will automatically be treated as a Shape.

That means we can treat different types in the same way, as long as they follow the Shape interface.

The smart-devices project teaches the following:
Create multiple interfaces that then use interface composition.
Polymorphism: one function UseDevices() could be used for many devices.