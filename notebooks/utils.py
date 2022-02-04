
import numpy as np
import tensorflow as tf
from tensorflow.keras.utils import get_custom_objects

def load_dataset(loc=""):
    def preprocess(x):
        x = x.astype(np.float32) / 255 # gray scale to floating point
        x = np.expand_dims(x, axis=3)
        return x
    x_train = np.load(loc + 'train_img.npy')
    x_train = preprocess(x_train)
    y_train = np.load(loc + 'train_labels.npy')
    
    x_val = np.load(loc + 'valid_img.npy')
    x_val= preprocess(x_val)
    
    y_val = np.load(loc + 'valid_labels.npy')
    
    x_test = np.load(loc + 'test_img.npy')
    x_test = preprocess(x_test)
    
    y_test = np.load(loc + 'test_labels.npy')
    return x_train, y_train, x_val, y_val, x_test, y_test


def loss(labels, logits):
    return tf.keras.losses.binary_crossentropy(labels, logits, from_logits=True)

def load_model(path="model/Logistic_Regression[0.986167][0.991453]-Mon Jan 31 10:15:08 2022.h5"):
    get_custom_objects().update({"loss": loss})
    model = tf.keras.models.load_model(path)
    return model 