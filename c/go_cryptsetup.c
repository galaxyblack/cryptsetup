#include <stdio.h>
#include <libcryptsetup.h>

void log_stdout_callback(int level, const char *msg, void *usrptr) {
	printf("%s", msg); fflush(stdout);
}

void go_cryptsetup_set_log_callback_to_stdout(struct crypt_device *cd) {
	crypt_set_log_callback(cd, log_stdout_callback, NULL);
}
