import re

txt = "lf AND lq -> ls"

x = re.compile("(.+) (.+) (.+)", txt)
y = re.compile("(.+?) -> (.+)", txt)

m = x.match(txt)