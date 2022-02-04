"""Setup for the HELR package"""

from distutils.errors import CompileError
from subprocess import call

from setuptools import Extension, setup
from setuptools.command.build_ext import build_ext
import os
class build_go_ext(build_ext):
    """Custom command to build the Golang extension from sources"""
    def build_extension(self, ext):
        ext_path = self.get_ext_fullpath(ext.name)
        cmd = ['go', 'build', '-buildmode=c-shared', '-o', ext_path]
        cmd += ext.sources 
        print(">>>>", ext.sources, ext_path)
        print("CWD:", os.getcwd())
        print("Command:", cmd)
        out = call(cmd)
        if out != 0:
            raise CompileError('Go build failed')

setup(
    name='HELR',
    version='0.1.0',
    description='A package for private Logistic Regression Inference',
    author='José Cabrero Holgueras',
    author_email='jcabreroholgueras@gmail.com',
    license='MIT',
    py_modules=['HELR'],
    ext_modules=[
        Extension('_simulator', ['cmd/export/main.go'])
    ],
    cmdclass={'build_ext': build_go_ext},
    install_requires=[
     'numpy',
    ],
    setup_requires=[
      'numpy',
    ],
    zip_safe=False
)
