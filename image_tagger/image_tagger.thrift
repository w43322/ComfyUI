namespace go comfyui.image_tagger

struct ProvinceMeta {
    1: i64 province_id;
    2: string province_name;
    3: string owner_tag;
    4: string province_religion;
    5: string province_culture;
    6: string province_culture_group_title;
}

struct Image {
    1: string url;
    2: i64 width;
    3: i64 height;
    4: list<string> tags;
    100: ProvinceMeta province_meta;
    101: i64 batch_index;
}

struct GetImagesReq {
    1: string tags (api.query="tags");
}

struct GetImagesResp {
    1: list<Image> images;
}

service ImageService {
    GetImagesResp GetImages(1: GetImagesReq request) (api.get="/get_images");
}
