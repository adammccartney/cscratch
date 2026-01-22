/* hello_sine.c */
#include <stdio.h>
#include <math.h>

#define SAMPLE_RATE 44100
#define NSECS 3
#define NSAMPLES (NSECS * SAMPLE_RATE)
#define PI 3.14159265
#define FREQ 442

int main(int argc, char* argv[])
{
    int i;
    float sample;
    for(i = 0; i < NSAMPLES; i++) {
        sample = sin(2 * PI * FREQ * i / SAMPLE_RATE);
        printf("%f\n", sample);
    }
}

