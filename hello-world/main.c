#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>

int main(int argc, char** argv) {
	printf("parent: %d\n", getpid());

	pid_t p = fork();

	if (p < 0) {
		printf("failed to fork\n");
		exit(1);
	}

	if (p == 0) {
		printf("child: %d\n", getpid());
	}

	sleep(20);	
}
