package editor;

/**
* Provider is the interface for all providers of content to the editor.
* Functions to work with generic provider should also be here.
*
* A provider is a source of content for a buffer, and should receive events
* from the buffer as input and interactions occur.
*
* IMPORTANT: events should be able to be emitted from other sources than direct
* user input. For example, a user navigating in another buffer of a file, should
* be able to trigger an event which the help_provider will receive, allowing it
* to update the help, without the user having to interact with the help buffer
* directly.
* 
* IDEA: User interactions should emit events, and the current buffer should pass
* event to it's provider, which should be able to emit events, which all other 
* active providers (and other systems) may be subscribed to. Any provider which
* receives an event should also be able to emit events. Maybe to avoid deep loops 
* or infinite loops, or just cascading explosions of events, each event should have
* a "depth" or "level" which is incremented each time the event is emitted, and it
* would be set to +1 of the event which triggered it. 
*
*
*
*/
