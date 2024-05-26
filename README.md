contexts : contexts of two things 1 . cancellation and progration

1 . cancellation : this helps the some process to cancel at any point of time.

2 . ctx : = context.Background() - background won't be cancelled , this is the default or root context.

3 . context type in context package is an interface , there are main four methods in it . 3 methods related to cancellation and propgation , and one method related to value propagation.

4. withvalue , withcancel , withdeadline are the child methods of contexts which are called on top of context.background . Context is a tree sturcture where Background is starting poing and its childs are withvalue , withcancel , withdeadline. without background child's can't be created.

5 . If the root context gets cancelled in a process , it passes to all its children and all the child contexts gets cancelled.

6. withcontext() - is used to set the context and context() is used to retrieve the context.

7. Context.withvalue(ctx , key , value) : it is a key value pair used to send with request contexts and response contexts.
