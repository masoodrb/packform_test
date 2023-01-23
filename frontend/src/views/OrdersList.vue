<template>
  <div class="container">
    <Header
      @onSearch="onSearch"
    >
    </Header>
    <div>
      <label>Created date</label>
    </div>
    <div>
      <el-date-picker 
        v-model="selectedDate" 
        type="daterange"
        range-separator="-"
        start-placeholder="Start date"
        end-placeholder="End date"
        @change="onDateChange"
      />
    </div>
    <div>
      Total Amount:
    </div>
    <DataGrid></DataGrid>
  </div>
</template>
<script lang="ts" setup>
import moment from 'moment'

import { ref, onMounted } from 'vue';
import { useOrderStore } from '@/stores/order'
import Header from '@/components/Header/index.vue'
import DataGrid from '@/components/DataGrid/index.vue'

const { loadOrderDetails } = useOrderStore()
onMounted(() => {
  loadOrderDetails()
})
const onSearch = (value: string) => {
  loadOrderDetails(value)
}

const selectedDate = ref([
  moment().startOf('month').format('YYYY-MM-DD'),
  moment().endOf('month').format('YYYY-MM-DD')
]);

const onDateChange = (value) => {
  console.log(moment(value[0]))
}

</script>

<style lang="scss" scoped>

.container {
  padding: 15px;
}
</style>
