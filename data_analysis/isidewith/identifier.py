import class_seperator, data_parser, common_qual

# Importing required variables
class_rep = class_seperator.class_rep 
res_dict = data_parser.RES_DICT


# Importing required functions
qual_common_res = common_qual.qual_common_res
form_parse = data_parser.form_parse
class_seperate = class_seperator.class_seperate


# Compile dict of all qn-ans pairs common to all forms for each class_pol (i.e. political party)
UID_DICT = {}


# Is a dict of all qn-ans pairs UNIQUE to each class_pol
FINAL_UID = {}
final_arr = ''


def find_qn_uid(in_path = './golden_reference', form_count = 32, out_path = './out_file.txt'):

	class_seperate(in_path)
	form_parse(form_count)

	# rep_arr_in : array of indices from class_rep for each class_pol
	# form_f : contains qn-res pairs for form f (reference)
	# form_f : contains qn-res pairs for form x (iterator)
	rep_arr_in = {}
	form_f = {}
	form_x = {}


	# uid_res : contains all qn-ans pairs unique to each class_pol
	uid_res = {}


	# To compare response to 'Quality'
	qual = {}
	q_dict = {}


	# f : gives the form index [0 - 31]
	f = 0
	x = 0
	q = 0


	# Iterating through each class_pol
	for class_pol in class_rep:

		q = 0
		flag = 0
		q_dict = {}

		uid_res = {}

		rep_arr_in = class_rep[class_pol]

		# Reference form = form at zeroth index in class_rep
		f = rep_arr_in[0]
		form_f = res_dict[f]

		# Iterating through all indices (other than reference)
		for i in range(len(rep_arr_in)-1):

			# Calc indices for each index (other than reference)
			x = rep_arr_in[i+1]
			form_x = res_dict[x]

			for qn in form_x:

				# Compiling a dict of all qualities for given class_pol
				if (str(form_x[qn])[0] == '['):
					
					q_dict[0] = form_f[qn]
					q_dict[q+1] = form_x[qn]
					q += 1

			# Comparing form_f, form_x only for the first time
			# Adding common qn-res pairs to uid_res
			if(flag == 0):

				for qn in form_f:

					if (str(form_x[qn])[0] != '['):

						qn_res_f = form_f[qn]['ans'].split(',')[0]
						qn_res_x = form_x[qn]['ans'].split(',')[0]
						
						if(qn_res_f == qn_res_x):					
							uid_res[qn] = qn_res_f

			# Comparing uid_res, form_x only for the first time
			# Retaining if qn-res pair common, deleting otherwise
			if(flag == -1):

				# Using temp to avoid runtime error
				# Here, uid_res is is manipulated before a loop iteration is complete
				# Throws error if uid_res is used in for-loop
				temp = uid_res.copy()
				for qn_uid in temp:

					if(uid_res[qn_uid] == form_x[qn_uid]['ans'].split(',')[0]):
						uid_res[qn_uid] = uid_res[qn_uid]

					else:
						del uid_res[qn_uid]

			flag = -1

		# Calc any common res to 'Qualities' header for givn class_pol
		qual = qual_common_res(q_dict)
		
		# Append to uid_res of given class_pol ONLY IF qual is not empty
		# Ignore qual otherwise
		if(len(qual) != 0):
			uid_res['Qualities'] = qual
			
		UID_DICT[class_pol] = uid_res

	find_unique(out_path)


def find_unique(out_path):

	# Picking out only the UNIQUE qn-ans pairs
	for pol_f in UID_DICT:

		temp = UID_DICT[pol_f].copy()
		res_f = UID_DICT[pol_f]

		# Looping over each qn for a given pol_f
		for qn in res_f:

			# Looping over each pol_x for a qn in a given pol_f
			# To eliminate any common pairs
			for pol_x in UID_DICT:

				res_x = UID_DICT[pol_x]

				# Ensure comparison if with all other political parties and not itself
				if(pol_f == pol_x):
					continue

				else:
					# If same qn serves as uid for pol_f and pol_x
					if((qn in res_x) == True):
						
						# If res to the common qn is the same for pol_x and pol_f
						# Remove qn from temp
						if((res_x[qn] == res_f[qn]) & (qn in temp)):
							del temp[qn]
					
					else:
						continue

		FINAL_UID[pol_f] = temp

	final_str = str(FINAL_UID).replace('}', '\n').replace('{', '\n').replace(", '", ",\n'").replace(',', '')

	final_arr = final_str.split(', ')
	write_to_file(out_path, final_arr)


# Function to over-write content of check_uid.txt
def write_to_file(file_path, final_arr):

	write_data = ''

	check_uid = open(file_path, 'w+')

	for ele in final_arr:
		write_data += ele

	check_uid.write(write_data)
	check_uid.close()
