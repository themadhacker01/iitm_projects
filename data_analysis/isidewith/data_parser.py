# Parsing through qn and heading sets
# Assuming that a JSON of headings and question are given

import json

# RES_DICT : dictionary of key-val pairs (here, qn-res) for a given form
RES_DICT = {}

# k = form no.
k = 0
count = 0

def form_parse(form_count):

	count = form_count

	# form_res = responses to ALL qns for given heading = {qn_res, ...}
	# qn_res = response to given qn
	# all_form_res = responses to ALL headings for given form_no
	all_form_res = {}
	form_res = {}
	qn_res = {}

	# Compiling response for header - 'Quality'
	# since it is in a different format
	qual_res = []
	head_qn = []

	qn = ''

	for k in range(count):
		
		form_file = open('./forms/' + str(k+1))
		data = json.loads(form_file.read())[k]

		form_res = {}

		# item : each header in given data
		for x, item in enumerate(data):		

			# inner_1 : list object of all qn-res sets for a given header
			inner_1 = data[item]		

			for	y, item_1 in enumerate(inner_1):

				# Reset vals for each different header
				qn_res = {}
				qual_res = []

				# item_1 : corresponds to qn for a given header
				qn = item_1	

				# inner_2 : response object for a given qn
				# res_type is list for headers (other than 'Quality')
				inner_2 = inner_1[item_1]		

				# For header 'Quality', res_type is array
				if(item == 'Qualities'):
					for q in range(len(inner_2)):
						qual_res.append(inner_2[q]['title'])			
				
				for z, item_2 in enumerate(inner_2):
					
					if(item == 'Qualities'):
						break

					else:
						inner_3 = inner_2[item_2]
						if((item_2 == 'question')):
							qn_res['ans'] = inner_3

						if((item_2 == 'importance')):
							qn_res['imp'] = inner_3
				
				form_res[qn] = qual_res if (item == 'Qualities') else qn_res

		RES_DICT[k] = form_res
