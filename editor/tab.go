// die/editor/tab.go

package editor;

/**
*  Tab is a container which holds a Layout, The layout may have multiple splits each
*  split will have a buffer. Where a buffer is a region of renderable text. And The
*  Buffer will have a `Source` which is the source of its contents. The source may
*  be a FileProvider, a LanguageServer, a Console, a LiveHelp Provider, etc.
*
*  In addition to a Split, a Window may be present in the Layout, which floats above
*  the underlying splits, and may contain a Buffer. Useful for things like a floating
*  Console popup, or a context help popup, auto complete
*
* EXAMPLE Editor State
* --------------------
*
* Tab 1 ^
*   Layout (Focused Split: RightTop - after window closed)
*     Split "Left"
*       Buffer (LiveHelp)  - for example a buffer which a live context help may render help for the symbol under the cursor of the active buffer
*     Split "RightTop" ^
*       Buffer (Current File)  - in this example the Active Split buffer, the current file being edited
*     Split "RightBottom"
*       Buffer (Console)  - in this example a buffer which is a console for the editor
*  Window ^   (command promt or other popup not related to the split+buffer active)
* 
* Tab 2
*  Layout (Focused Split: Main)
*    Split "Main"
*      Buffer (Current File)  - in this example the Active Split buffer, the current file being edited
*
* --------------------
*
* Here there are multiple Tabs open, one with a multi-function layout with a console, 
* help, and file all open, there is a temporary popup window that launched to type a command. 
* The other Tab is a simple layout with only a single split and a single buffer.
*
* Maybe it should be allowed to have multiple buffers stacked in a split, but only one
* being active at a time. Hopefully making it easy to switch between multiple open buffers
* with a single hotkey, or even a Buffer Switcher popup window.
*/

import (
    "fmt"
)

type Tab struct {
    
}

func GetTabAbout() string {
    return fmt.Sprintf("die %s\n", "Tab About")
}


