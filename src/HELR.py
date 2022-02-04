import ctypes
#from numpy.ctypeslib import ndpointer
import numpy as np
import time


from distutils.sysconfig import get_config_var
from pathlib import Path

# Location of shared library
here = Path(__file__).absolute().parent
ext_suffix = get_config_var('EXT_SUFFIX')
so_file = here / ('_simulator' + ext_suffix)

class HELR(object):
    __MAX_SHAPE = 128 * 128
    def __init__(self, weights, bias):
        if weights.shape[0] < HELR.__MAX_SHAPE:
            raise Exception("The expected shape of the LR is 128 x 128")
        
        # Adapting weights as parameters
        self.w = weights.flatten()
        self.w = (ctypes.c_float *  (self.w.shape[0]))(*self.w)
        self.b = ctypes.c_float(bias)

        so = ctypes.cdll.LoadLibrary(so_file)

        self.test = so.test
        self.test.argtypes = [ctypes.c_int, ctypes.c_float * (128 * 128), # img shape + img ptr
                             ctypes.c_int, ctypes.c_float * (128 * 128), # weights shape + weights ptr
                             ctypes.c_float, # bias
                             ctypes.c_int] # reuse_keys
        self.test.restype = ctypes.c_double

        self.free = so.free
        self.free.argtypes = [ctypes.c_void_p]

    def predict(self, img, reuse_keys=False):
        x = img.flatten()
        shape = x.shape[0]
        xpp = (ctypes.c_float *  shape)(*x)
        c_reuse_keys = 1 if reuse_keys else 0
        #toc = time.time()
        ret = self.test(ctypes.c_int(shape), xpp, len(self.w), self.w, self.b, c_reuse_keys)
        #tic = time.time()
        #print("GO TIME: ", toc - tic)
        return ret

    def get_lr_params(tf_model):
        """This function takes a tensorflow logistic regression and adapts it.
        It assumes a sigmoid function as output function
        """
        weights = tf_model.layers[1].weights[0].numpy()
        bias = tf_model.layers[1].weights[1].numpy()
        return weights, bias


