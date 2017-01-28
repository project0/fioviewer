<template>
  <div>
    <h2 class="graphTitle">#{{index}} <a class="btn btn-secondary" href="#" role="button" data-toggle="modal"
                                         :data-target="'#modal-' + _uid">Edit</a> <span class="badge badge-default">{{title}}</span>
    </h2>
    <chart :chartData="chartData" :options="options" ref="myChart"></chart>


    <!-- Modal -->
    <div class="modal fade" :id="'modal-' + _uid" tabindex="-1" role="dialog">
      <div class="modal-dialog modal-lg" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <span class="badge badge-info">#{{index}}</span>
            <input class="form-control" type="text" v-model="title">

          </div>
          <div class="modal-body">
            <!-- edit section -->

            <h4>Type</h4>
            <div class="form-check">
              <label class="form-check-label">
                <input class="form-check-input" type="radio" value="bw" :name="'graph_type_' + _uid" v-model="type">
                Bandwith
              </label>
              <label class="form-check-label">
                <input class="form-check-input" type="radio" value="iops" :name="'graph_type_' + _uid" v-model="type">
                IOPS
              </label>
              <label class="form-check-label">
                <input class="form-check-input" type="radio" value="lat" :name="'graph_type_' + _uid" v-model="type">
                Latency
              </label>
            </div>

            Start offset: <input class="form-control" type="text" v-model="range.start">
            <hr>
            <div v-for="n, i in files">
              <select class="custom-select" v-model.trim="files[i]">
                <option :value="f.filename" v-for="f, idx in getList(type)">{{f.name}} ({{f.type.name}})</option>
              </select>
            </div>

            <div>
              <a class="btn btn-primary" href="#" role="button" @click.prevent="files.push('')">Add Log</a>
            </div>

          </div>
          <div class="modal-footer">
            <button type="button" role="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
            <button type="button" role="button" class="btn btn-success" data-dismiss="modal" @click="updateGraph">
              Update
            </button>
          </div>
        </div>
      </div>
    </div>

  </div>

</template>
<script>
  import LineChart from './chartjs/LineChart'

  export default {
    name: 'graph',
    props: {
      index: {
        type: Number,
        required: true
      },
      width: {
        default: '100%'
      },
      height: {
        default: '400px'
      }
    },
    data () {
      return {
        chartData: {},
        options: {
          scales: {
            yAxes: [
              {
                ticks: {
                  callback: function (label, index, labels) {
                    return label
                  }
                },
                scaleLabel: {
                  display: true,
                  labelString: ''
                }
              }
            ]
          }
        },
        type: 'bw',
        files: [],
        list: [],
        range: {
          max: 0,
          start: 0,
          end: 0
        },
        title: 'New chart ' + this.index
      }
    },
    components: {
      chart: LineChart
    },
    methods: {
      fetchList: function (e) {
        this.$http.get('list').then((response) => {
          this.list = response.body
        })
      },
      getList: function (filter) {
        if (filter) {
          return this.list.filter(function (v) {
            return v.type.group === filter
          })
        }
        return this.list
      },
      updateGraph: function (e) {
        this.files = this.files.filter(function (v) {
          return v !== ''
        })

        this.$http.post('graph', {
          maxDataPoints: Math.round(window.innerWidth / 4),
          range: {
            start: Number(this.range.start),
            end: Number(this.range.end)
          },
          aggregation: 'avg',
          files: this.files
        }).then((response) => {
          for (var i in response.body.datasets) {
            response.body.datasets[i].fill = false
            response.body.datasets[i].lineTension = 0
            response.body.datasets[i].borderCapStyle = 'butt'
            response.body.datasets[i].pointBorderWidth = 1
            response.body.datasets[i].pointHoverRadius = 5
            response.body.datasets[i].pointRadius = 1
            response.body.datasets[i].pointHitRadius = 10
          }
          // set yaxes labelstring
          this.options.scales.yAxes[0].scaleLabel.labelString = response.body.logs[0].type.name + ' - ' + response.body.logs[0].type.unit
          this.chartData = response.body
        })
      }
    },
    created () {
      this.fetchList()
    }
  }

</script>

<style>
  .graphTitle {

  }
</style>
