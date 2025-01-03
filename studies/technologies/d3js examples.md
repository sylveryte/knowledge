# d3js examples

## Basic bar chart with animation and upate

```typescriptreact
import { useEffect, useState } from "react";
import * as d3 from "d3";

function getRandoInts(x = 50) {
  return Array(x)
    .fill(1)
    .map((_) => Math.round(Math.random() * 50));
}

function App() {
  const [d, setD] = useState(getRandoInts());
  const h = 400;
  const w = 800;

  const pd = 20;
  const ah = h - pd;
  const yScale = d3.scaleLinear().domain([0, 50]).range([ah, 0]);
  const xScale = d3
  .scaleBand()
  .domain(d3.range(d.length))
  .range([0, w])
  .round(true)
  .paddingInner(0.5);
  const svg = d3.select("svg");
  const maxd = d3.max(d) || ah;

  useEffect(() => {
    svg
      .selectAll("rect")
      .data(d)
      .enter()
      .append("rect")
      .attr("height", (d) => ah - yScale(d))
      .attr("width", xScale.bandwidth())
      .attr("fill", (dx, i) =>
        i % 5 == 0
          ? `rgba(4,255,25,${dx / maxd})`
          : `rgba(255,0,0,${dx / maxd})`,
      )
      .attr("x", (dx) => `${xScale(dx)}`)
      .attr("y", (d) => yScale(d));

    svg
      .selectAll("line")
      .data(d)
      .enter()
      .append("line")
      .attr("stroke", (_, i) => (i % 5 == 0 ? "green" : "red"))
      .attr("stroke-width", (_, i) => (i % 5 == 0 ? "1" : "0"))
      .attr("x1", (dx) => `${xScale(dx)}`)
      .attr("y1", (d) => yScale(d))
      .attr("x2", pd)
      .attr("y2", (d) => yScale(d));

    svg
      .selectAll("text")
      .data(d)
      .enter()
      .append("text")
      .attr("fill", "yellow")
      .text((d) => d)
      .attr("x", (dx) => `${xScale(dx)}`)
      .attr("y", (d) => pd - 5 + yScale(d));

    const ay = d3.axisLeft(yScale);
    const ax = d3.axisBottom(xScale);
    svg.append("g").call(ay).attr("transform", `translate(${pd},${pd})`);
    svg
      .append("g")
      .call(ax)
      .attr("transform", `translate(0,${h - pd})`);
  }, []);

  useEffect(() => {
    const svg = d3.select("svg");
    svg
      .selectAll("rect")
      .data(d)
      .transition()
      .duration(2000)
      .attr("fill", (dx, i) =>
        i % 5 == 0
          ? `rgba(4,255,25,${dx / maxd})`
          : `rgba(255,0,0,${dx / maxd})`,
      )
      .attr("x", (dx) => `${xScale(dx)}`)
      .attr("y", (d) => pd - 20 + yScale(d))
      .attr("height", (d) => ah - yScale(d));

    svg
      .selectAll("line")
      .data(d)
      .transition()
      .duration(2000)
      .attr("x1", (dx) => `${xScale(dx)}`)
      .attr("y1", (d) => pd + yScale(d))
      .attr("x2", pd)
      .attr("y2", (d) => pd + yScale(d));

    svg
      .selectAll("text")
      .data(d)
      .transition()
      .text((d) => d)
      .duration(2000)
      .attr("y", (d) => pd - 5 + yScale(d))
      .attr("x", (dx) => `${xScale(dx)}`);
  }, [d]);

  const onUpdate = () => {
    setD(getRandoInts());
  };

  return (
    <main>
      <h1>D3 js</h1>
      <svg height={h} width={w}></svg>
      <button onClick={() => onUpdate()}>Update</button>
    </main>
  );
}

export default App;
```

## Basic bar chart with label and yaxis with line indicating projecting on yscale

```typescriptreact
import { useEffect, useState } from "react";
import * as d3 from "d3";

function App() {
const [d] = useState(
Array(50)
.fill(1)
.map((_) => Math.round(Math.random() * 50)),
);
const h = 400;
const w = 800;

useEffect(() => {
const svg = d3.select("svg");
const pd = 20;
const ah = h - pd;
const bw = (w - pd) / d.length;
const maxd = d3.max(d) || ah;
const linscalex = d3.scaleLinear().domain([0, 50]).range([ah, 0]);

svg
.selectAll("rect")
.data(d)
.enter()
.append("rect")
.attr("height", (d) => ah - linscalex(d))
.attr("width", bw)
.attr("fill", (dx, i) =>
i % 5 == 0
? `rgba(4,255,25,${dx / maxd})`
: `rgba(255,0,0,${dx / maxd})`,
)
.attr("x", (_, i) => pd + i * bw)
.attr("y", (d) => pd + linscalex(d));

svg
.selectAll("line")
.data(d)
.enter()
.append("line")
.attr("stroke", (_, i) => (i % 5 == 0 ? "green" : "red"))
.attr("stroke-width", (_, i) => (i % 5 == 0 ? "2" : "1"))
.attr("x1", (_, i) => pd + i * bw)
.attr("y1", (d) => pd + linscalex(d))
.attr("x2", pd)
.attr("y2", (d) => pd + linscalex(d));

svg
.selectAll("text")
.data(d)
.enter()
.append("text")
.attr("fill", "yellow")
.text((d) => d)
.attr("x", (_, i) => 5 + pd + i * bw)
.attr("y", (d) => pd - 5 + linscalex(d));

const ax = d3.axisLeft(linscalex);
svg.append("g").call(ax).attr("transform", `translate(${pd},${pd})`);
}, [d]);

return (
<main>
<h1>D3 js</h1>
<svg height={h} width={w}></svg>
</main>
);
}

export default App;
```

## Basic bar chart with label on top

```typescriptreact
function App() {
  const [d] = useState(
    Array(50)
      .fill(1)
      .map((_) => Math.round(Math.random() * 50) + 5),
  );
  const h = 200;
  const w = 800;
  const pd = 5;
  const ah = 200 - 4 * pd;
  const bw = (w - 2 * pd) / d.length;
  const linscalex = d3.scaleLinear().domain([0, 50]).range([0, ah]);

  useEffect(() => {
    d3.select("svg")
      .selectAll("rect")
      .data(d)
      .enter()
      .append("rect")
      .attr("height", (d) => linscalex(d))
      .attr("width", bw)
      .attr("fill", (d) => `rgba(255,0,0,${linscalex(d) / ah})`)
      .attr("x", (_, i) => pd + i * bw)
      .attr("y", (d) => 2 * pd + ah - linscalex(d));

    d3.select("svg")
      .selectAll("text")
      .data(d)
      .enter()
      .append("text")
      .attr("fill", "yellow")
      .text((d) => d)
      .attr("x", (_, i) => pd + i * bw)
      .attr("y", (d) => ah + 24 - linscalex(d));
  }, [d]);

  return (
    <main>
      <h1>D3 js</h1>
      <svg height={h} width={w}></svg>
    </main>
  );
}
```

---
