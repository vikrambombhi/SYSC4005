import util
import numpy as np

ws1 = util.get_data("data/ws1.dat")
ws2 = util.get_data("data/ws2.dat")
ws3 = util.get_data("data/ws3.dat")

util.hist(ws1, "ws1 hist")
util.hist(ws2, "ws2 hist")
util.hist(ws3, "ws3 hist")

util.create_plots(ws1, "weibull_min", "ws1")
util.create_plots(ws2, "weibull_min", "ws2")
util.create_plots(ws3, "weibull_min", "ws3")

m = np.mean(ws1)
util.qqplot(ws1, "skewnorm", (m), "ws1_skewnorm")
m = np.mean(ws2)
util.qqplot(ws2, "skewnorm", (m), "ws2_skewnorm")
m = np.mean(ws3)
util.qqplot(ws3, "skewnorm", (m), "ws3_skewnorm")
