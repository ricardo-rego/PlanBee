# -*- coding: UTF-8 -*-

import zen
import json

with open(r".\v6.json", "r") as f:
  content = f.read()

engine = zen.ZenEngine()
decision = engine.create_decision(content)

data = {
   "documento":"35077857000194",
   "nsu":1,
   "tipoPessoa": "PJ",
   "tipoPagamento": "A",
   "trxPos":1500.00,
   "trxAcum":10000,
   "trxMes":5000000.00,
   "trxHorario":"1330",
   "renda":90000,
   "faturamento":100000,
   "cnae":3313901,
   "porte":"ME",
   "mccRisco":"F",
   "csnu":"F",
   "fronteira":"F"
}

result = decision.evaluate(data)
print(json.dumps(result, indent=4))

'''
{
    "result": [
        {
            "alerta": {
                "idregra": "PLDFT_C1B_UIF1073_PorteME",
                "alerta": true
            }
        },
        {
            "alerta": {
                "idregra": "PLDFT_TDB_C1B_UIF1045_MovMes",
                "alerta": true
            }
        },
        {
            "alerta": {
                "idregra": "PLDFT_TDB_C1B_UIF1073_PorteME",
                "alerta": true
            }
        }
    ],
    "performance": "17.4467ms"
}
'''