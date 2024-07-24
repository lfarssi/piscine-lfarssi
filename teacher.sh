#!/bin/bash

# Step 1: Find the key interview number
key_interview=$(grep -oP '(?<=export interview_id=)\d+' interviews/interview-*)

# Step 2: Store the interview number in an environment variable
export KEY_INTERVIEW=$key_interview

# Step 3: Print the environment variable containing the interview number
echo "Key Interview Number: $KEY_INTERVIEW"

# Step 4: Print the content of the interview
echo "Content of the Key Interview:"
cat interviews/interview-$KEY_INTERVIEW

# Step 5: Print the content of the MAIN_SUSPECT environment variable
echo "Main Suspect: $MAIN_SUSPECT"
