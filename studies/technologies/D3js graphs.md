# D3 graphs

{https://d3js.org/}[link]

# Learn Basics

## Bind to data

- d3 can manipulate htmls as well as svgs
- bind _data_ and create all the elements
- below code creates circles inside svg with data by state d
  - note we selectAll before even elements exists,
    it will either replace or create new circles
  - `data(d)` binds the data to process,
    it makes it looping possible for each element
  - most value like `fill` below can take fx for conditional value
  - `enter()` is what enables the creation of new elements

```typescriptreact
const [d] = useState(
 Array(5)
   .fill(1)
   .map((_) => Math.round(Math.random() * 25) + 25),
   );
 const h = 200;
 const w = 200;
 const linscalex = d3
   .scaleLinear()
   .domain([25, 50])
   .range([5, h / 6]);

 useEffect(() => {
   d3.select("svg")
   .selectAll("circle")
   .data(d)
   .enter()
   .append("circle")
   .attr("r", (d) => linscalex(d))
   .attr("stroke", "pink")
   .attr("stroke-width", 3)
   .attr("fill", (d) => (d % 2 ? "red" : "indigo"))
   .attr("cx", (dx, i) => linscalex(dx) + i * linscalex(w / d.length))
   .attr("cy", linscalex(h / 2));
   <svg height={h} width={w}></svg>
```

## d3 goodies

### range

    - `d3.range(5) //[0,1,2,3,4]`

### min max

    - `d3.min(dataset, (d)=>d.value)`
    - `d3.max(dataset, (d)=>d.value)`

## Scale

- provided scales in d3js are just functions with state
- _clever thing_ to invert y axis simply _invert_ the range
- _domain_ (your data) == /scale fx/ ==>> _range_ (chart values)
- update scale with `let newscale = scaleInstance.domain/range(newvalus)`
- eg: {\* Examples }

### Updating scale

    - easy `yscale.domain([new values])`

### more methods

- `nice()` makes rounding for nicer labels,
  eg 0.93484 => 1
- `round(true)`
- `rangeRound()` nearest whole nuber
- `clamp()` forces all outputs to be within range,
  excessive value will be rounded to low/high of range
- has padding builtin
  `.paddingInner(0.5)` /not tried/

### Other scales

#### `scaleSqrt` square root

#### `scalePow` to the power of

#### `scaleLog`

#### `scaleQuantize` linear scale with discrete values,

     for sorted bucket

#### `scaleQuantile` same above but for domain input

#### `scaleOrdinal` for nonquantitative values like name

#### `scaleBand`

- for many visuals(bars etc),
  that are in arbitary order,
  but must be evenly spaced
- divides into even bands based on length of the input domain

#### `scaleTime` _for time_

```typescriptreact
xScale = d3.scaleTime().domain([Date Date]).range(0,400)
```

## Axis

- Axis can be added for bottom,left,right,top
- Axis take scale to make easily all data
- you have to transform most of the time
- can give how many ticks, it's approximate merely suggestion
  or exact ticks values
- can format tick labels
- eg {:$/studies/technologies/d3js examples:\*\* Basic bar chart with label and yaxis with line indicating projecting on yscale}[Basic bar chart with label and yaxis with line indicating projecting on yscale]
- eg

```typescriptreact
const ax = d3.axisLeft(linscalex);
// or
const ax = d3.axisLeft().scale(linscalex);

//draw
svg.append("g").call(ax).attr("transform", `translate(${x},${y})`);
```

## Update & transition

- no `append` or `enter`
  eg: {\*\* Basic bar chart with animation and upate}

```typescriptreact
svg
  .selectAll("line")
  .data(d)
  .transition()
  .duration(2000)
  .attr("x1", (dx) => `${xScale(dx)}`)
```

- add transition delay, duration n more
- can have _staggering_ transition
  `.delay((_, i) => i * 100)` after transition
- _no_ two transition can run parallelly,
  d3 only runs one of them(latest)
- can have on _start_ on _end_ transition /(didnt get it to work)/

```typescriptreact
.transition()
.duration(4000)
.on(
"start",
() =>
function () {
d3.active(this).attr("fill", "white").transition();
},
)
```

- eg {\*\* Basic bar chart with animation and upate}

```typescriptreact
svg
.selectAll("line")
.data(d)
.transition()
.duration(2000)
.attr("x1", (dx) => `${xScale(dx)}`)
```

### Can do multiple transition

in one go searially

```typescriptreact
.data(d)
.transition() // transision 1 *
.duraiont....
.on("start",()=>{.....})
.attrsss....
.transition() // transition2
.duration...
.attr
```

### Named transition

- to do multiple transisions
  `.transition("somename)`
- Interrupt
  `.selectAll('').interrupt("somename)`

### Contain with clipping

- Steps
- 1.  Define clippath and give it an ID

```typescriptreact
svg.append("clippath")
.attr("id","chart-area")
.append("rect") // within clippath
.attrs... // rect boundry is clippath boundry ??
```

- 2.  Put visuals elements within clippath
- 3.  Add a ref to clippath from,
      whatever elements you wish to be masked,
      to put all many elements within clippath
      we'll add all elements in a group element `g`
      and add ref to g

```typescriptreact
svg.append("g")
.attr("id","circles") //good idea to give groups id
.attr("clippath","url(#chart-area)")
.selectAll("circle")
.data(d)....contineue adding circles
```

### Add one `append`

- `data()` gives ref of all the elements we have binded to,
  thats how we call `update`
  `const bars = svg.selectAll('rects').data(dataset)` all rect refs
- we have added one value to dateset, we can use `enter()`
  to address new corresponding dom element without touching existing rects
  `bars.enter().append('rect').....attrs` all attrs except width
  since width will not be calulated corerctly and will be just on right off,
  since above `enter` gives only one selection to new element

### Merge new and existings

- We will merge old and new existing ones for smooth transition
  continue with above

```typescriptreact
.merge(bars)  //existings
.transition()
...rest attrs
```

### Remove one `exit`

- `bars.exit()` gives elements without data
- to smoothly remove

```typescriptreact
bars.exit()
.transition()
.attr("x",w)
.remove() //special method waits for transition to complete
```

- `remove()` waits for transition to end then removes from DOM

### Data joins with keys

- default join is `index` as key
- we can use `key functions` to control the join
- with data `.data(dataset,d=>d.key)`

## Interactivity

### on

- on for event listening
- eg `d3.select("p").on("click",()=>doSomething())`

## hover

- use css easy

## Grouping svgs

- check `g element` in {\*\*\* Contain with clipping}

## Tooltip

### Default

- ugly browser one
  `.append('title').text('soms')`

### Svg element

- complicated but doable
- get x,y from event function
  `.on("mouseover",function (d){const x = d3.select(this).attr('x)})`

### Make a div

- to make a tooltip best way is to make a div,
  hide and on hover change top left and fill data

## Cick to Sort

- easy

```typescriptreact
.sort((a,b)=>d3.ascending(a,b))
```

## Consider touch devices

- checkout d3.touch

* Learn some advance d3

* Case study

## Gantt chart

- {:$/accrete/cassini/work_items/Gantt chart building:}[Gantt chart building]

## Force Directed Graph, Disjoint

one that obsidian uses
{https://observablehq.com/@d3/disjoint-force-directed-graph}[link]

## Force directed graph

{https://observablehq.com/@d3/force-directed-graph-component}[link]

## Force directed tree

{https://observablehq.com/@d3/force-directed-tree?intent=fork}[link]

- Learn Advance

## Using Paths

### Line chart

### Dealing with missing data

### Area charts

## Selections

### Close look at Selections

    more specific filtering n all

### Storing selections

### Merging selections

### Filtering selections

### To each() Their own

    `selection.each((d,i)=>some action with this ref to element)`

## Layouts

### List of layouts availble in d3js

- Chord
- Cluster
- Force
- Pack
- Partition
- Pie
- Stack
- Tree
- Treemap

### Pie layout

### Stack layoutS

### Stacked Areas

### Force layout

- relates to {:$/projects/neorgraph:}[Neorgraph]

#### Defining the force simulations

#### Creating visual elements

#### Updating visuals over time

#### Draggable nodes

## Geomapping

### JSON, meet GeoJSON

- specific json for geo data

### Path

### Projections

### Choropleth

### Adding points

### Panning

### Transitioning the Map

### Dragging the Map

### Border Problems

### Zooming

- `d3.zoom` isnt just for geographic maps,
  can be used for _panning_ and _zooming_ for any chart type

#### Constrainng panning and zooming

## Exporting

### Bitmaps

### PDF

### SVG

## Project walk through

- Case studies

  - few case studies in the book, can also look on d3js site

- Examples

  - {:$/studies/technologies/d3js examples:\*\* Basic bar chart with animation and upate}[Basic bar chart with animation and upate]
  - {:$/studies/technologies/d3js examples:\*\* Basic bar chart with label and yaxis with line indicating projecting on yscale}[Basic bar chart with label and yaxis with line indicating projecting on yscale]
  - {:$/studies/technologies/d3js examples:\*\* Basic bar chart with label on top}[Basic bar chart with label on top]

  ===

---

Origin: {:$/book/Interactive Data Visualization for the Web, 2nd edition 2017:}[Interactive Data Visualization for the Web, 2nd edition 2017]
