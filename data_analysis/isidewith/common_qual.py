# If common response to 'Quality' exists
# Retain ele_q
# And insert into uid_res{}
def qual_common_res(q_dict):

	qn_res_f = q_dict[0]
	flag = 0
	qual = {}
	q = 0

	for i in range(len(q_dict)-1):

		qn_res_x = q_dict[i+1]

		# Finding common ele between first 2 arrays of 'Qualities'
		if(flag == 0):
			
			flag = -1
			for item_f in qn_res_f:
				for item_x in qn_res_x:
					
					if(item_f == item_x):
						qual[q] = item_f
						q += 1

		else:

			# Using ret{} to determine which items to delete /retain
			# Initialising ret{} with -1
			ret = {}
			for j in qual:
				ret[j] = -1

			# -1 : delete, 0: retain
			for q in qual:
				for item_x in qn_res_x:				
					if(qual[q] == item_x):
						ret[q] = 0

			itr = qual.copy()
			
			# Remove item from dict qual if no common vals exist
			for ele in itr:
				if(ret[ele] == -1):
					del qual[ele]

	return qual