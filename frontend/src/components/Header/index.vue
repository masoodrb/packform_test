<template>
  <el-row>
    <el-col :md="3" :sm="6" :lg="3">
      <div class="customHeader">
        <div class="searchLabel">
          <font-awesome-icon
            icon="fa-solid fa-magnifying-glass"
            class="fa-search-icon"
            size="3x"
          />
          <label>Search</label>
        </div>
      </div>
    </el-col>
    <el-col :sm="18" :md="20">
      <el-input
        v-model="query"
        placeholder="Search Here"
        @change="$emit('onSearch', query)"
      >
      </el-input>
    </el-col>
  </el-row>
</template>
<script setup lang="ts">
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { ref } from "vue";

const query = ref("");
</script>
<style lang="scss">
.searchLabel {
  font-weight: 600;

  label {
    font-size: 2em;
    font-weight: bold;
  }

  .fa-search-icon {
    width: 30px;
    height: 30px;
    margin-left: 8px;
    margin-right: 8px;
  }
}
</style>
