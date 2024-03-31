import logging
import sys
from comfy_api_simplified import ComfyApiWrapper, ComfyWorkflowWrapper
from custom_nodes.province_selector.info_fetcher import province_strs, province_map

logging.basicConfig(stream=sys.stdout, level=logging.INFO)

api = ComfyApiWrapper()

wf = ComfyWorkflowWrapper("comfy_api_simplified/examples/workflow_api.json")



for s in province_strs:
    i = int(s.split(",")[0])
    province_map[i] = s

# ok = set([
#     # 1
#     201, 401, 601, 701, 2201, 2301, 2401, 3001, 4101, 4501, 4701, 4801,
#     1, 101, 901, 1201, 4201,
#     4301, 301,
#     2101, 
#     # 2
#           ])

# idx = 10
batch = 0

for idx in range(1, 11, 1):
    for i in range(idx, 10000, 100):
        if i in province_map: #  and i not in ok:
            wf.set_node_param("EU4 Province Selector", "province", province_map[i])
            wf.set_node_param("EU4 Province Selector", "seed_increment", batch)
            api.queue_prompt(wf)

# wf.save_to_file("modified_wf.json")
