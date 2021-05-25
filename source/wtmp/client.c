#include <stdlib.h>

#include "common/common.h"
#include "libgocli.h"

int main(int argc, char* argv[]) {
	// For command-line arguments
	struct Arguments args;

	/*
	 * Flag to determine whether or not to
	 * do busy-waiting and non-blocking calls
	 * TODO we have elided this
	 * busy_waiting = check_flag("busy", argc, argv);
	 */

	parse_arguments(&args, argc, argv);

	ClientMain(&args);

	return EXIT_SUCCESS;
}
