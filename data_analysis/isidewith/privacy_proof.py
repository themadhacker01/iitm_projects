import os, sys
import identifier

find_qn_uid = identifier.find_qn_uid

ip_path = ''
op_path = ''
form_count = 0

########################## take input parameters ################################

def exit_with_help(error=''):
	print("""\
Usage: identifier.py [options]

options:
	-inpath : 	Path to the file containing form no. with associated political parties (default = golden_reference)
	-count : 	No. of forms to consider in the 'forms' folder (default = 32)
	-outpath : 	Output file path (default = ./out_file.txt)
""")
	
	print(error)
	sys.exit(1)

# Arguments to be read from command line
args = [('inpath', 'inpath', 'inpath'),
		('count', 'count', 'count'),
		('outpath', 'outpath', 'outpath')]

# Checking if all variables are/will be set
for var, env, arg in args:

	if not '-'+arg in sys.argv:		
		vars()[var] = os.getenv(env)		

		if ((vars()[var] == None) & (var != 'outpath')):
			exit_with_help('Error: Environmental Variables or Argument'+' insufficiently set! ("-'+arg+'")')

# Read parameters from command line call
if len(sys.argv) != 0:
	
	i = 0
	options = sys.argv[1:]
	
	# iterate through parameters
	while i < len(options):
		
		if options[i] == '-inpath':
			i = i + 1
			ip_path = options[i]
		
		elif options[i] == '-count':
			i = i + 1
			form_count = int(options[i])
		
		elif options[i] == '-outpath':
			i = i + 1
			op_path = options[i]

		else:
			i = i + 1

if((ip_path == '') | (form_count == 0)):
	exit_with_help('Error: Unknown Argument! ('+ options[i] + ')')

if(op_path == ''):
	op_path = './out_file.txt'
	find_qn_uid(ip_path, form_count, op_path)

else:
	find_qn_uid(ip_path, form_count, op_path)