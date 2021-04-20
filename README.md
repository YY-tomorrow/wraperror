dao 层遇到sql.ErrNoRows应该warp并且抛到上层处理，

service 层中有多种操作，包括查询，删除，如果查询是报错sql.ErrNoRows，应用withMessage，补充信息抛出给上层，如果delete操作，先获取信息校验是否可以删除，如果报错sql.ErrNoRows，说明数据已经不存在，将错误降级返回。

业务层报错答应错误日志。
            