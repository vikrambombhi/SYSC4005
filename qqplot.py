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
    filepath = title+".png"
    pylab.savefig(filepath)
    pylab.show()


def hist(measurements, title):
    num_bins = math.sqrt(len(measurements))
    num_bins = int(num_bins)
    hist, bins = np.histogram(measurements, bins=num_bins)
    width = 0.7 * (bins[1] - bins[0])
    center = (bins[:-1] + bins[1:]) / 2
    pylab.title(title)
    pylab.bar(center, hist, align='center', width=width)
    pylab.show()


ws1 = get_data("data/ws1.dat")
ws2 = get_data("data/ws2.dat")
ws3 = get_data("data/ws3.dat")


hist(ws1, "ws1 hist")
qqplot(ws1, "frechet_r", (1.5), "ws1 frechet_r 1.5")
