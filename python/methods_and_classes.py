from __future__ import print_function


def test_method():
	print("Hello world!")


class A(object):

	def __init__(self, some_var):
		self._semiprivate = some_var
		self._name = "objecta"
		print('Initializing with', self._semiprivate)

	def __str__(self):
		return self._name

	def some_method(self, x):
		return x * self._semiprivate

	@staticmethod
	def static_method():
		print("Hello im a static method")

	@classmethod
	def class_method(cls):
		print("Hello im class method")

if __name__ == "__main__":
	a = A(15)
	x = ("Hello there, {0} ".format(a))
	print(x) # will print out Hello there, objecta
	print(a.some_method(5))
	test_method() # prints Hello world
	a.static_method()
	A.class_method()
else:
	print("Help, I'm being imported")