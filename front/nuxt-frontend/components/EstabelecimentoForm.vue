<template>
  <form
    @submit.prevent="emitirSubmit"
    class="bg-white p-6 rounded-xl max-w-4xl mx-auto space-y-6"
  >
    <h2 class="text-xl font-semibold text-gray-800">Informações do Estabelecimento</h2>

    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <div>
        <label class="block text-sm font-medium text-gray-700">Nome</label>
        <input
          v-model="form.nome"
          type="text"
          class="mt-1 block w-full border border-gray-300 rounded-md p-2 focus:ring-blue-500 focus:border-blue-500"
          required
        />
      </div>

      <div>
        <label class="block text-sm font-medium text-gray-700">Razão Social</label>
        <input
          v-model="form.razao_social"
          type="text"
          class="mt-1 block w-full border border-gray-300 rounded-md p-2 focus:ring-blue-500 focus:border-blue-500"
          required
        />
      </div>

      <div>
        <label class="block text-sm font-medium text-gray-700">Endereço</label>
        <input
          v-model="form.endereco"
          type="text"
          class="mt-1 block w-full border border-gray-300 rounded-md p-2 focus:ring-blue-500 focus:border-blue-500"
          required
        />
      </div>

      <div>
        <label class="block text-sm font-medium text-gray-700">Número</label>
        <input
          v-model="form.numero"
          type="text"
          class="mt-1 block w-full border border-gray-300 rounded-md p-2 focus:ring-blue-500 focus:border-blue-500"
          required
        />
      </div>

      <div>
        <label class="block text-sm font-medium text-gray-700">Cidade</label>
        <input
          v-model="form.cidade"
          type="text"
          class="mt-1 block w-full border border-gray-300 rounded-md p-2 focus:ring-blue-500 focus:border-blue-500"
          required
        />
      </div>

      <div>
        <label class="block text-sm font-medium text-gray-700">Estado</label>
        <input
          v-model="form.estado"
          type="text"
          maxlength="2"
          class="mt-1 block w-full border border-gray-300 rounded-md p-2 focus:ring-blue-500 focus:border-blue-500"
          required
        />
      </div>

      <div>
        <label class="block text-sm font-medium text-gray-700">CEP</label>
        <input
          v-model="form.cep"
          type="text"
          class="mt-1 block w-full border border-gray-300 rounded-md p-2 focus:ring-blue-500 focus:border-blue-500"
          required
        />
      </div>

      <div>
        <label class="block text-sm font-medium text-gray-700">Número do Estabelecimento</label>
        <input
          v-model="form.numero_estabelecimento"
          type="text"
          class="mt-1 block w-full border border-gray-300 rounded-md p-2 focus:ring-blue-500 focus:border-blue-500"
          required
        />
      </div>
    </div>

    <div class="flex justify-end">
      <button
        type="submit"
        class="bg-blue-600 text-white px-6 py-2 rounded-md hover:bg-blue-700 transition duration-200"
      >
        Salvar
      </button>
    </div>
  </form>
</template>

<script setup lang="ts">
import { reactive, watch } from 'vue'

interface Estabelecimento {
  id?: number
  nome: string
  razao_social: string
  endereco: string
  cidade: string
  estado: string
  cep: string
  numero: string
  numero_estabelecimento: string
}

const props = defineProps<{ estabelecimento: Estabelecimento }>()

const emit = defineEmits<{
  (e: 'submit', payload: Estabelecimento): void
}>()

const form = reactive<Estabelecimento>({ ...props.estabelecimento })

watch(
  () => props.estabelecimento,
  (novo) => {
    if (novo) Object.assign(form, novo)
  },
  { immediate: true, deep: true }
)

function emitirSubmit() {
  emit('submit', JSON.parse(JSON.stringify(form)))
}
</script>
