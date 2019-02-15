import numpy as np
import pylab
import scipy.stats as stats
import math
import pdb

def get_data(filename):
    with open(filename) as f:
        content = f.readlines()

    return [float(x.strip()) for x in content]


def qqplot(measurements, dist, sparams, title):
    stats.probplot(measurements, dist=dist, sparams=sparams, plot=pylab)
    pylab.title(title)
    filepath = "figures/"+title+".png"
    pylab.savefig(filepath)
    # pylab.show()
    pylab.close()


def hist(measurements, title):
    num_bins = math.sqrt(len(measurements))
    num_bins = int(num_bins)
    hist, bins = np.histogram(measurements, bins=num_bins)
    width = 0.7 * (bins[1] - bins[0])
    center = (bins[:-1] + bins[1:]) / 2
    pylab.title(title)
    pylab.bar(center, hist, align='center', width=width)
    filepath = "figures/"+title+".png"
    pylab.savefig(filepath)
    # pylab.show()
    pylab.close()

def create_plots(measurements, plot_type, title):
    i=0.5
    while (i<=5):
        t = title+"_"+plot_type+"_"+str(i)
        qqplot(measurements, plot_type, (i), t)
        i+=0.5
