import gc
import torch
import numpy as np
import torchvision
import torchvision.transforms as transforms
import PIL.Image as Image

import proto.ml_pb2 as struct

CATEGORY_NAMES = [
    '__background__', 'person', 'bicycle', 'car', 'motorcycle', 'airplane', 'bus',
    'train', 'truck', 'boat', 'traffic light', 'fire hydrant', 'N/A', 'stop sign',
    'parking meter', 'bench', 'bird', 'cat', 'dog', 'horse', 'sheep', 'cow',
    'elephant', 'bear', 'zebra', 'giraffe', 'N/A', 'backpack', 'umbrella', 'N/A', 'N/A',
    'handbag', 'tie', 'suitcase', 'frisbee', 'skis', 'snowboard', 'sports ball',
    'kite', 'baseball bat', 'baseball glove', 'skateboard', 'surfboard', 'tennis racket',
    'bottle', 'N/A', 'wine glass', 'cup', 'fork', 'knife', 'spoon', 'bowl',
    'banana', 'apple', 'sandwich', 'orange', 'broccoli', 'carrot', 'hot dog', 'pizza',
    'donut', 'cake', 'chair', 'couch', 'potted plant', 'bed', 'N/A', 'dining table',
    'N/A', 'N/A', 'toilet', 'N/A', 'tv', 'laptop', 'mouse', 'remote', 'keyboard', 'cell phone',
    'microwave', 'oven', 'toaster', 'sink', 'refrigerator', 'N/A', 'book',
    'clock', 'vase', 'scissors', 'teddy bear', 'hair drier', 'toothbrush'
]


class CarDetector:
    def __init__(self, threshold):
        self.device = torch.device("cuda:0" if torch.cuda.is_available() else "cpu")
        self.model = torchvision.models.detection.maskrcnn_resnet50_fpn(pretrained=True)
        self.model.eval()
        self.loader = transforms.Compose([transforms.ToTensor()])
        self.threshold = threshold

    def predict(self, image):
        image = Image.fromarray(np.uint8(image)).convert('RGB')
        image = self.loader(image)
        image = image.to(self.device)
        with torch.no_grad():
            pred = self.model([image])
        pred_score = list(pred[0]['scores'].detach().numpy())
        pred_t = [pred_score.index(x) for x in pred_score if x > self.threshold][-1]
        masks = (pred[0]['masks'] > 0.5).squeeze().detach().cpu().numpy()
        pred_class = [CATEGORY_NAMES[i] for i in list(pred[0]['labels'].numpy())]
        pred_boxes = [[(i[0], i[1]), (i[2], i[3])] for i in list(pred[0]['boxes'].detach().numpy())]
        masks = masks[:pred_t + 1]
        pred_boxes = pred_boxes[:pred_t + 1]
        pred_class = pred_class[:pred_t + 1]
        result = {val: struct.Rows() for val in {'bicycle', 'car', 'motorcycle', 'bus', 'truck'}}
        for i in range(len(masks)):
            if result.get(pred_class[i]):
                f_box = struct.Boxes(x=pred_boxes[i][0][0], y=pred_boxes[i][0][1])
                s_box = struct.Boxes(x=pred_boxes[i][1][0], y=pred_boxes[i][1][1])
                row = struct.Row()
                row.area = np.sum(masks[i])
                row.boxes.extend([f_box, s_box])
                result[pred_class[i]].data.extend([row])
        del image, pred
        gc.collect()
        return result
