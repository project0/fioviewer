<template xmlns:v-on="http://www.w3.org/1999/xhtml">
  <div id="app">
    <div class="container">
      <a class="btn btn-primary" href="#" role="button" @click.prevent="addGraph">Add Chart</a>
    </div>
    <div>
      <div class="graphContainer" v-for="props, index in graphs">
        <graph :index="index" :style="{ height: props.height, width: props.width }" class="container"></graph>
      </div>
    </div>
  </div>
</template>

<script>
  import Graph from './components/Graph'

  export default {
    name: 'app',
    components: {
      graph: Graph
    },
    data () {
      return {
        graphs: []
      }
    },
    methods: {
      addGraph: function (e) {
        this.graphs.push({
          width: '100%',
          height: '500px'
        })
      }
    },
    created () {
      this.addGraph()
    }
  }
</script>

<style>
  .graphContainer {
    padding-bottom: 50px;
  }
</style>
