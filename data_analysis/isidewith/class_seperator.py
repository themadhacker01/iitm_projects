# To seperate given set of classes (political parties) from their serial nos.

import json

# Compiling the classes as a dictionary
CLASS_DICT = {}

# rep_arr : array of all indices of a given class_pol
# class_rep : list of rep_arr for ALL class_pol
rep_arr = []
class_rep = {}

def class_seperate(in_path):

	file_path = in_path

	# Opening the file golden_reference
	class_file = open(file_path, 'r')
	strArr = class_file.readlines()

	for line in strArr:
		
		line = line.replace('\n', '')
		br = line.index(' ')
		i = int(line[0:br])
		CLASS_DICT[i-1] = line[br:].strip()

	for i in CLASS_DICT:
		
		# Reset vals for each class_pol
		ctr = 0
		rep_arr = []

		# Iterates through the list to check for equality
		for j in CLASS_DICT:
					
			if(CLASS_DICT[i] == CLASS_DICT[j]):
				ctr = (ctr+1)
				rep_arr.append(j)

		class_rep[CLASS_DICT[i]] = rep_arr
