package editor;

/**
* Script Provider, is really just a provider which interfaces with the Editor, and allows
* a developer/user to create their own Provider via a script. This can be to integrate with
* external tools, or to create custom functionality directly in the script. It should be
* able to make all of the same Event Subscriptions, and Emit any Event just as any other
* Provider. And should be able to send contents to the Buffer as any other Native Provider.
*
* This should be extra useful for integrating with external Language Servers.
*
*
*
*/
