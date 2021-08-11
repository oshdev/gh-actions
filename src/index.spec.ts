import { tag } from "./tag";

describe("tagging", () => {
  it("tags", () => {
    expect(tag("osh")).toEqual("osh was here");
  });
});
