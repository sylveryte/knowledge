# <%tp.file.title%>

<%* if (tp.date.now("d","P-1D",tp.file.title,"YYYY-MM-DD dddd") == 0) { %> [[<%tp.date.now("YYYY-MM-DD dddd","P-3D",tp.file.title,"YYYY-MM-DD dddd")%>]] <%* } else { %> [[<%tp.date.now("YYYY-MM-DD dddd","P-1D",tp.file.title,"YYYY-MM-DD dddd")%>]] <%* } %> <== <%tp.file.title%> ==>  <%* if (tp.date.now("d","P+1D",tp.file.title,"YYYY-MM-DD dddd") == 6) { %> [[<%tp.date.now("YYYY-MM-DD dddd","P+3D",tp.file.title,"YYYY-MM-DD dddd")%>]] <%* } else { %> [[<%tp.date.now("YYYY-MM-DD dddd","P+1D",tp.file.title,"YYYY-MM-DD dddd")%>]]
<%* } %>

---

## Tasks
- [ ] <%tp.file.cursor(2)%>

---
Sprint: [[<%tp.file.cursor(1)%>]]