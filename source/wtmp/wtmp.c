#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/wait.h>

#include "common/parent.h"
#include "common/process.h"

int main(int argc, char* argv[]) {
    char* me_path = getenv("WTMP_ME_PATH");
    if (me_path == NULL) {
        printf("You need to set WTMP_ME_PATH\n");
        exit(1);
    }

	pid_t c1_id = start_child(me_path, argc, argv);
	sleep(1);

	setup_parent("wtmp", argc, argv);

	waitpid(c1_id, NULL, WUNTRACED);
}
