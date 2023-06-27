Singleton design pattern – having a unique
instance of a type in the entire program

The Singleton pattern is easy to remember. As the name implies, it will provide you with a
single instance of an object, and guarantee that there are no duplicates.
At the first call to use the instance, it is created and then reused between all the parts in the
application that need to use that particular behavior.

You'll use the Singleton pattern in many different situations. For example:
When you want to use the same connection to a database to make every query
When you open a Secure Shell (SSH) connection to a server to do a few tasks,
and don't want to reopen the connection for each task
If you need to limit the access to some variable or space, you use a Singleton as
the door to this variable (we'll see in the following chapters that this is more
achievable in Go using channels anyway)
If you need to limit the number of calls to some places, you create a Singleton
instance to make the calls in the accepted window

As a general guide, we consider using the Singleton pattern when the following rule
applies:
We need a single, shared value, of some particular type.
We need to restrict object creation of some type to a single unit along the entire
program.

From: Go design patterns. Mario Castro Contreras. Copyright © 2017 Packt Publishing

