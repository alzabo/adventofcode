const std = @import("std");

pub fn main() !void {
    var file = try std.fs.cwd().openFile("input.txt", .{});
    defer file.close();

    var buf_reader = std.io.bufferedReader(file.reader());
    var in_stream = buf_reader.reader();

    var buf: [1024]u8 = undefined;
    var last: u32 = 0;
    var ct: u32 = 0;

    while (try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        const v = std.fmt.parseInt(u32, line, 10);
        // TODO: Review unwrapping error unions.
        // Not real confident that these are being handled as well as they could be
        if (v) |val| {
            if (last > 0 and val > last) {
                ct += 1;
                std.log.info("{d} gt {d}", .{ val, last });
            }
            last = val;
        } else |_| {
            unreachable;
        }
        //std.log.info("line {s}", .{line});
    }
    std.log.info("ct: {}", .{ct});
}
