<template lang="">
 <el-row>
      <el-table :data="tableData" :default-sort="{ prop: 'date', order: 'descending' }" style="width: 100%">
        <el-table-column prop="orderName" label="Order name" />
        <el-table-column prop="companyName" label="Customer Company" />
        <el-table-column prop="customerName" label="Customer name" />
        <el-table-column prop="orderDate" label="Order date" sortable :formatter="formatter" />
        <el-table-column prop="deliveredAmount" label="Delivered amount" />
        <el-table-column prop="totalAmount" label="Total Amount" />
      </el-table>
    </el-row>
    <el-row>
      <div class="pagination-wrapper">
        <el-pagination :current-page="1" :page-size="5" :page-sizes="[5, 10]"
          layout="total, sizes, prev, pager, next, jumper" 
          :total="16" 
          @size-change="() => { }"
          @current-change="() => { }" />
      </div>
    </el-row>
</template>
<script lang="ts" setup>

import type { TableColumnCtx } from 'element-plus'
import type { OrderDetails } from '@/api'


const tableData = Array<OrderDetails>();

const formatter = (row: OrderDetails, column: TableColumnCtx<OrderDetails>) => {
  return row.orderDate
}

</script>
<style lang="scss">
.pagination-wrapper {
  width: 100%;
  margin-top: 20px;

  .el-pagination {
    justify-content: center;
  }
}
</style>