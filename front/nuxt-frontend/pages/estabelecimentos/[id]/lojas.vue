<template>
  <div class="p-6 max-w-4xl mx-auto">
    <h1 class="text-2xl font-bold mb-4">
      Lojas do Estabelecimento: {{ nomeEstabelecimento }}
    </h1>

    <div v-if="lojas.length > 0" class="space-y-4">
      <div
        v-for="loja in lojas"
        :key="loja.id"
        class="border rounded p-4 shadow hover:shadow-md transition"
      >
        <h2 class="text-lg font-semibold">{{ loja.nome }}</h2>
        <p class="text-sm text-gray-600">Número: {{ loja.numero_loja }}</p>
        <p class="text-sm text-gray-600">
          Endereço: {{ loja.endereco }}, {{ loja.numero }}
        </p>
        <p class="text-sm text-gray-600">
          Cidade: {{ loja.cidade }} - {{ loja.estado }}
        </p>

        <div class="flex gap-2 mt-2">
          <NuxtLink
            :to="`/lojas/${loja.id}`"
            class="bg-yellow-500 hover:bg-yellow-600 text-white px-3 py-1 rounded text-sm"
          >
            Editar
          </NuxtLink>

          <button
            @click="removerLoja(loja.id)"
            class="bg-red-600 hover:bg-red-700 text-white px-3 py-1 rounded text-sm"
          >
            Remover
          </button>
        </div>
      </div>
    </div>

    <p v-else class="text-gray-500">Nenhuma loja encontrada para este estabelecimento.</p>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from '#app'

interface Loja {
  id: number
  numero_loja: string
  nome: string
  razao_social: string
  endereco: string
  cidade: string
  estado: string
  cep: string
  numero: string
  estabelecimento_id: number
}

interface Estabelecimento {
  id: number
  nome: string
}

const route = useRoute()
const { $api } = useNuxtApp()

const lojas = ref<Loja[]>([])
const nomeEstabelecimento = ref('')

const carregarEstabelecimento = async () => {
  try {
    const id = route.params.id
    const response = await $api<{ data: Estabelecimento }>(`/estabelecimentos/${id}`)
    nomeEstabelecimento.value = response.data.nome
  } catch (error) {
    console.error('Erro ao carregar nome do estabelecimento', error)
  }
}

const carregarLojas = async () => {
  try {
    const id = route.params.id
    const response = await $api<{ data: Loja[] }>(`/estabelecimentos/${id}/lojas`)
    lojas.value = response.data
  } catch (error) {
    console.error('Erro ao carregar lojas', error)
    alert('Erro ao carregar lojas')
  }
}

const removerLoja = async (id: number) => {
  const confirmar = confirm('Tem certeza que deseja remover esta loja?')
  if (!confirmar) return

  try {
    await $api(`/lojas/${id}`, {
      method: 'DELETE'
    })
    lojas.value = lojas.value.filter((loja) => loja.id !== id)
    alert('Loja removida com sucesso!')
  } catch (error) {
    console.error('Erro ao remover loja', error)
    alert('Erro ao remover loja')
  }
}

onMounted(async () => {
  await carregarEstabelecimento()
  await carregarLojas()
})
</script>
