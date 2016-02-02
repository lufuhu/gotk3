// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12

// See: https://developer.gnome.org/gtk3/3.14/api-index-3-14.html

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import "unsafe"

// PrefersAppMenu is a wrapper around gtk_application_prefers_app_menu().
func (v *Application) PrefersAppMenu() bool {
	return gobool(C.gtk_application_prefers_app_menu(v.native()))
}

// GetActionsForAccel is a wrapper around gtk_application_get_actions_for_accel().
func (v *Application) GetActionsForAccel(acc string) []string {
	cstr1 := (*C.gchar)(C.CString(acc))
	defer C.free(unsafe.Pointer(cstr1))

	var acts []string
	c := C.gtk_application_get_actions_for_accel(v.native(), cstr1)
	originalc := c
	defer C.g_strfreev(originalc)

	for *c != nil {
		acts = append(acts, C.GoString((*C.char)(*c)))
		c = C.next_gcharptr(c)
	}

	return acts
}
